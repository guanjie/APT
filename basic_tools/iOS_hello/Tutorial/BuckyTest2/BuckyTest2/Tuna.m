//
//  Tuna.m
//  BuckyTest2
//
//  Created by humancool on 4/9/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import "Tuna.h"

static int gx = 0; // only makes it global not external

@implementation Tuna
-(void) addOne{
    gx++;
    num++;
}
-(void) printIt{
    NSLog(@"The number is %i", gx);
}

@end
