//
//  tutorial.m
//  tutorial
    //
//  Created by humancool on 3/31/15.
//  Copyright (c) 2015 humancool. All rights reserved.
//

#import <Foundation/Foundation.h>

@interface Person : NSObject{
    int age;
    int weight;
}
-(void) print;
-(void) setAge: (int) a;
-(void) setWeight: (int) w;
-(int) getAge;
-(int) getWeight;

@end

@implementation Person
-(void) print{
    NSLog(@"I am %i years old and weigh %i pounds", age, weight);
}

//-(void) setAge: (int) a{
//    age = a;
//}
//
//-(void) setWeight: (int) w{
//    weight = w;
//}
//
//-(int) getAge{
//    return age;
//}
//
//-(int) getWeight{
//    return weight;
//}

@end

int main(int argc, char *argv[]){
    @autoreleasepool {

//        int time;
//        
//        int count = 1;
//        
//        do {
//            count++;
//            
//            NSLog(@"Enter the fn time!");
//            scanf("%i", &time);
//            
//            if (time < 11){
//                NSLog(@"Good morning to ya!");
//            } else if (time < 16) {
//                NSLog(@"Godo afternoon to ya!");
//            } else if ( time <= 24) {
//                NSLog(@"Godo night hoss!");
//            } else {
//                NSLog(@"WTF did you even enter?");
//            }
//        } while (count <= 10);
        
        int a = 2;
        int b = 2;
        int c = 4;
        
        a == b? NSLog(@"they are equal"): NSLog(@"they are different");
        c? NSLog(@"True"): NSLog(@"False");
        
        return 0;
    }
}




