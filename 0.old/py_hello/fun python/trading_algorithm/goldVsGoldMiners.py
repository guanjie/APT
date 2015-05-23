{rtf1ansiansicpg1252cocoartf1138cocoasubrtf510
{fonttblf0fswissfcharset0 Helvetica;}
{colortbl;red255green255blue255;}
margl1440margr1440vieww15620viewh17040viewkind0
deftab720
pardtx720tx1440tx2160tx2880tx3600tx4320tx5040tx5760tx6480tx7200tx7920tx8640pardeftab720pardirnatural

f0fs36 cf0 from zipline.transforms.utils import EventWindow
from zipline.transforms.stddev import MovingStandardDevWindow
from zipline.transforms.mavg import MovingAverageEventWindow
from zipline.utils.protocol_utils import ndict

import statsmodels.api as sm

class OLSWindow(EventWindow):
    def __init__(self, window_size, stock1, stock2):
        # Call the superclass constructor to set up base EventWindow
        # infrastructure. This call will create a window of window_size trading days -
        # holidays and weekends will be excluded from the tally.
        EventWindow.__init__(self, True, window_size, None)
        self.stock1 = stock1
        self.stock2 = stock2

        self.last_calc = None
        self.daycount = window_size
        self.refresh_period = window_size // 10
        # flag for signaling that the event window is full
        # i.e. we have window_size trading days of data
        self.full = False
        self.beta = None
        self.zscore = None
        self.spread_stddev = MovingStandardDevWindow(True, self.refresh_period, None)
        self.spread_mavg = MovingAverageEventWindow(['spread'], True, self.refresh_period, None)

    def handle_data(self, data):
        """
        New method to handle a data frame as sent to the algorithm's handle_data
        method.
        """

        price1 = data[self.stock1].price
        price2 = data[self.stock2].price
        spread = price1 - price2
        dt = data[self.stock1].datetime
        if dt < data[self.stock2].datetime:
            dt = data[self.stock2].datetime

        event = ndict({
                  'dt'    : data[self.stock1].datetime,
                  'price1': data[self.stock1].price,
                  'price2': data[self.stock2].price,
                  'spread': spread
                })

        self.update(event)

        if not self.beta:
            return

        current_date = data[self.stock1].datetime
        self.spread = data[self.stock1].price - self.beta * data[self.stock2].price
        # at the moment, the standard deviation event window only looks at a price field,
        # while moving average will take a parameter for the field to average. So, our
        # event fakes out the std deviation event window by repeating the spread in the price.
        spread_event = ndict({
                          'price' : self.spread,
                          'spread': self.spread,
                          'dt' : current_date
                          })
        # calculate the stddev of the spread
        self.spread_stddev.update(spread_event)
        # calculate the average of the spread
        self.spread_mavg.update(spread_event)

        spread_avg = self.spread_mavg.get_averages()['spread']
        spread_std = self.spread_stddev.get_stddev()
        if spread_std:
            self.zscore = (self.spread - spread_avg)/spread_std

    def handle_add(self, event):
        """
        This method is required by EventWindow, and is called for
        each new event added to the window. There isn't an incremental
        implementation for OLS available (clone and create one! We'll send you a t-shirt!), 
        so rather than recalculate every minute, we hold a value for refresh_period days
        and then recalculate.
        """
        if not self.last_calc:
            self.last_calc = event.dt
            return

        ols_age = event.dt - self.last_calc
        if ols_age.days >= self.refresh_period:
            # EventWindow holds the current window of data in a list
            # called ticks. Here we are splitting it into two lists of
            # prices (floats) to pass to the Ordinary Least Squares (OLS)
            # implemented in statsmodels.


            p1 = [x.price1 for x in self.ticks]
            p2 = [x.price2 for x in self.ticks]

            model = sm.OLS(p1, p2)
            results = model.fit()
            self.beta = results.params[0]
            self.last_calc = event.dt

    def handle_remove(self, event):
        """
        This method is required by EventWindow, and is called whenever
        an event falls out of the active window. Because the OLS implementation
        we are using from statsmodels is not iterative, and is quite heavy computationally,
        we ignore this event and only periodically re-calcuate the OLS in handle_add (above)
        """
        # since an event is expiring, we know the window is full
        self.full = True


def initialize(context):
    context.gld = sid(26807)
    context.gdx = sid(32133)
    # there are 252 trading days per year on US Markets, so
    # 126 days is a 6 month window.
    # 6-month rolling ols calculation, recalculated every 30 days
    # ols is making a linear fit between price of gld and price of gdx.
    context.ols_window = OLSWindow(126, context.gld, context.gdx)
    # calculate the std dev over


    # maximum total exposure (longs - shorts) in $US allowed
    context.max_notional = 1 * 1000 * 1000.0 #1M

def handle_data(context, data):
    context.ols_window.handle_data(data)
    if not context.ols_window.full:
        return

    # calculate the current notional value of each position
    notional1 = context.portfolio.positions[context.gld].amount * data[context.gld].price
    notional2 = context.portfolio.positions[context.gdx].amount * data[context.gdx].price

    # if notional1 is zero, we don't have a bet on and we can buy or sell
    if notional1 == 0:
        can_buy = True
        can_sell = True
        spread_bet = 0.0
    else:
        # check that our spread bet has an absolute exposure within our max_notional limit.
        spread_bet = abs(notional1) + abs(notional2) * notional1/abs(notional1)
        can_buy = spread_bet < 0 or spread_bet < context.max_notional
        can_sell = spread_bet > 0 or spread_bet > -1 * context.max_notional


    # hit the escape hatch if we don't have enough data to do calculations.
    zscore = context.ols_window.zscore
    beta = context.ols_window.beta

    bet_shares = 5
    if zscore >= 2.0 and can_sell:
        # sell the spread, betting it will narrow since it is over 2 std deviations 
        # away from the average
        order(context.gld, -1 * bet_shares)
        order(context.gdx, bet_shares * beta)

    elif zscore <= -2.0 and can_buy:
        # buy the spread
        order(context.gld, bet_shares)
        order(context.gdx, -1 * beta * bet_shares)

    elif zscore <= 1.0 and zscore >= -1.0:
        reduce_position(context.portfolio, context.gld, data, bet_shares)
        reduce_position(context.portfolio, context.gdx, data, bet_shares)


def reduce_position(portfolio, stock, data, abs_quantity):
    """
    decrease exposure, regardless of position long/short.
    buy for a short position, sell for a long.
    """
    pos_amount = portfolio.positions[stock].amount
    if pos_amount > 0:
        order(stock, -1 * abs_quantity)
    elif pos_amount < 0:
        order(stock, abs_quantity)
}
