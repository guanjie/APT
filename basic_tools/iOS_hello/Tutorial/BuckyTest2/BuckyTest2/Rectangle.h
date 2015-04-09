//
//  Rectangle.h
//  BuckyTest2
//
//  Created by humancool on 4/7/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import <Foundation/Foundation.h>

@class XYPoint;
@interface Rectangle : NSObject{
    int width;
    int height;
    XYPoint *origin;
}
@property int width, height; //define properties
-(XYPoint *) origin;
-(void) setOrigin:(XYPoint *)pt;
-(int) area;
-(int) perimeter;
-(void) setWH:(int) w: (int) h;

@end