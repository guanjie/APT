//
//  main.m
//  BuckyTest2
//
//  Created by humancool on 4/7/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import <Foundation/Foundation.h>
#import "Rectangle.h"
#import "XYPoint.h"
#import "Tuna.h"
#define EVEN(x) x%2==0

int gDrunk = 21;

int main(int argc, const char * argv[]) {
    @autoreleasepool {
        
        enum day {m=1, t=2, w=3, h=4, f=5};
//        enum day today;
        
        if (EVEN(3)) {
            NSLog(@"This is even");
        } else {
            NSLog(@"This id ODD");
        }
    }
    return 0;
}
