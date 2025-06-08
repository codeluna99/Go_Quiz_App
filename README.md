Thing to Do in the Project:
- Learn to read CSV file.
- Learn to pass command line arguments.
- Basic structure of Go Project.

Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Add On: 
    - The quiz should move to the next question if the user is not able to answer within the limit.
    - The skipped question should mark to zero.
    - Wrong answers should hold a penalty of 0.5 marks.
    - Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. Hint: Check out the strings package.
    - Add an option (a new flag) to shuffle the quiz order each time it is run.