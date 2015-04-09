//
//  Person.h
//  BuckyTest2
//
//  Created by humancool on 4/7/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import <Foundation/Foundation.h>

@interface Person : NSObject{
    int age;
    int weight;
}
@property int age, weight;
-(void) print;
-(void) dateAge: (int) a:(int) i;

@end
