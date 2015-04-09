//
//  Son.m
//  BuckyTest2
//
//  Created by humancool on 4/9/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import "Son.h"

@implementation Son : Mom
-(void) setNum1{
    num1 = 14;
    newVal = 69;
}
-(void) printNumber{
    NSLog(@"The number is %i %i ", num1, newVal);
}

@end
