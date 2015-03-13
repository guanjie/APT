import smtplib


def send(emailTos, subject, text):

    # emailAddress, emailPassword are string type
    emailAddress = "humancool@hotmail.com"
    emailPassword = "betterman"
    emailTo = ', '.join(emailTos)

    message = """\
From: %s
To: %s
Subject: %s

%s
""" % (emailAddress, emailTo, subject, text)

    print 'SENDING EMAIL:'
    print ''
    print ''
    print message
    server = smtplib.SMTP('locahost')  # port 465 or 587 'smtp.gmail.com'
    server.ehlo()
    server.starttls()
    server.ehlo()
    server.login(emailAddress, emailPassword)
    server.sendmail(emailAddress, emailTo, message)
    server.close()
    print ''
    print ''
    print 'SENT'


def main():
    emailTo = ['humancool@hotmail.com']
    subject = 'This is Subject'
    text = 'This is the text area'
    send(emailTo, subject, text)

if __name__ == "__main__":
    main()
