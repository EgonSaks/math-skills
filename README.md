# Objectives

The purpose of this project is for you to calculate the following:

- Average
- Median
- Variance
- Standard Deviation

# Instructions

Your program must be able to read from a file and print the result of each statistic mentioned above. In other words, your program must be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:
```
189
113
121
114
145
110
...
```
This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

To run your program a command similar to this one will be used if your project is made in Go:

```
$> go run your-program.go data.txt
```

After reading the file, your program must execute each of the calculations asked above and print the results in the following manner (the following numbers are only examples):

```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

Please note that the values are rounded integers.

# Testing

Your program will be tested by an auditor who will run a program provided by us. This program creates a random data set of numbers and prints the result. The auditor job is to compare how your program performed.

You can choose the language in which you want to build your program.

This project will help you learn about:

- Statistics and Mathematics
  - Average
  - Median
  - Variance
  - Standard Deviation

# How to test:

## Windows or Linux

Download [math-skills](https://assets.01-edu.org/stats-projects/math-skills) file and make sure it is in the same directory where you have `main.go` file. Give the math-skills file executable permissions:

```
chmod +x math-skills
```

Run the script with `./math-skills`, in current working directory it generates `data.txt` file.

```
./math-skills
```
## Mac OS

On `Mac OS` you need to have docker installed to run math-skills file while math-skills file is compiled in another environment.

```
docker run -v $PWD:/program/ -it --entrypoint sh alpine
```
```
cd /program
```
```
./math-skills
```
```
exit
```
After you have been getting `data.txt` run the `main.go` program in terminal with the created file `(data.txt)` by the previous command.

```
go run .
```
or
```
go run main.go
```

# What to look:

+ Are the outputs of both programs (the one provided and the student one) in the same format?

+ In the output of the student program, are the data types of the values rounded integers?

+ Did the values of both programs match?

Do the same procedure (running the script provided and the student program) 3 more times in order to test new data sets.

+ Did the values of both programs match in all tries?

# Happy coding :)



