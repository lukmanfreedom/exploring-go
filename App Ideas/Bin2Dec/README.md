# Bin2Dec

Program that convert binary number to decimal number.

## Description

Project ideas from [Bin2Dec](https://github.com/florinpop17/app-ideas/blob/master/Projects/1-Beginner/Bin2Dec-App.md).

This program will ask binary number input with maximum 8 binary digits, then will convert it to decimal number.

## Example

Input binary num = 10 => decimal number 2
-   (0 * 2^0) + (1 * 2^1) 
    <br>= (0 * 1) + (1 * 2) 
    <br>= 0 + 2 
    <br>= 2

Input binary num = 101 => decimal number 5
-   (1 * 2^0) + (0 * 2^1) + (1 * 2^2)
    <br>= (1 * 1) + (0 * 2) + (1 * 4)
    <br>= 1 + 0 + 4
    <br>= 5

## Challenges

-   Arrays may not be used to contain the binary digits entered by the user
-   Determining the decimal equivalent of a particular binary digit in the
    sequence must be calculated using a single mathematical function, for
    example the natural logarithm. It's up to you to figure out which function
    to use.

## User Stories

-   [v] User can enter up to 8 binary digits in one input field
-   [v] User must be notified if anything other than a 0 or 1 was entered
-   [v] User views the results in a single output field containing the decimal (base 10) equivalent of the binary number that was entered

## Bonus features

-   [v] User can enter a variable number of binary digits

## Approach to solve problem

-   Define a function that accepts an input.
-   Validate input (number only, accept only binary digit 0 & 1, max 8 digits).
-   Start for loop until input becomes 0.
-   Find the remainder of input and devide by 10.
-   Calculate decimal number using previous decimal number and remainder * pow(2, index).
-   Return decimal number.



