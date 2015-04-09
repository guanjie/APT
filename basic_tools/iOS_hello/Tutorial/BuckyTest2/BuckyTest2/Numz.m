//
//  Numz.m
//  BuckyTest2
//
//  Created by humancool on 4/9/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import "Numz.h"


@implementation Numz
-(void) setNumbers: (int) a: (int) b{
    num1 = a;
    num2 = b;
}
-(void) add{
    answer = num1 + num2;
}
-(void) print{
    NSLog(@"I am from the NUMBERS class biotch!, %i", answer);
}

@end
