//
//  Charz.m
//  BuckyTest2
//
//  Created by humancool on 4/9/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import "Charz.h"

@implementation Charz
-(void) setCharz{
    c1 = 'B';
    c2 = 'W';
}
-(void) add{
    NSLog(@"%c%c", c1, c2);
}
-(void) print{
    NSLog(@"I am from the charz class mofo.");
}

@end
