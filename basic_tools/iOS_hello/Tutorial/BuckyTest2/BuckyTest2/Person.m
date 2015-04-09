//
//  Person.m
//  BuckyTest2
//
//  Created by humancool on 4/7/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import "Person.h"

@implementation Person
@synthesize age,weight;
-(void) print{
    NSLog(@"I am %i years old and weigh %i pounds", age, weight);
}

-(void) dateAge: (int) a:(int) i {
    NSLog(@"You can date chicks %i years old", (a/2+7)-(i/10000));
}

@end
