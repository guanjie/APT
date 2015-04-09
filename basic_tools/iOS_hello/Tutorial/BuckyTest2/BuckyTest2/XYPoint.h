//
//  XYPoint.h
//  BuckyTest2
//
//  Created by humancool on 4/8/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import <Foundation/Foundation.h>

@interface XYPoint : NSObject{
    int x;
    int y;
}
@property int x, y;
-(void) setXY: (int) xVal: (int) yVal;

@end
