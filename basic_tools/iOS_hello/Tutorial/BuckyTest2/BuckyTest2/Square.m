//
//  Square.m
//  BuckyTest2
//
//  Created by humancool on 4/8/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import "Square.h"

@implementation Square : Rectangle
-(void) setSide: (int)s{
    [self setWH: s : s];
}
-(int) side{
    return width;
}

@end