//
//  Rectangle.m
//  BuckyTest2
//
//  Created by humancool on 4/7/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import "Rectangle.h"

@implementation Rectangle
@synthesize width, height; // setters and getters.
-(void) setWH:(int) w: (int) h{
    width = w;
    height = h;
}
-(int) area{
    return width * height;
}
-(int) perimeter{
    return (width + height)*2;
}
-(XYPoint *) origin{
    return origin;
}
-(void) setOrigin:(XYPoint *)pt{
    origin = pt;
}

@end
