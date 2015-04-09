//
//  Numz.h
//  BuckyTest2
//
//  Created by humancool on 4/9/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import <Foundation/Foundation.h>

@interface Numz : NSObject{
    int num1;
    int num2;
    int answer;
}
-(void) setNumbers: (int) a: (int) b;
-(void) add;
-(void) print;

@end
