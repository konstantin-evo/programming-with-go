<p align="center">
  <img src="src/go-logo.png" alt="Go Logo" width="30%">
</p>

<h1 align="center">Programming with Google Go Specialization</h1>

This repository contains the code for the "Programming with Google Go Specialization" on Coursera. The specialization
consists of several modules, each focusing on different aspects of programming in Go.

Specialization: [Programming with Google Go Specialization](https://www.coursera.org/specializations/google-golang)

Modules:

1. [Getting Started with Go](https://www.coursera.org/learn/golang-getting-started?specialization=google-golang)
2. [Functions, Methods, and Interfaces in Go](https://www.coursera.org/learn/golang-functions-methods?specialization=google-golang)
3. [Concurrency in Go](https://www.coursera.org/learn/golang-concurrency?specialization=google-golang)

## Getting Started with Go

The module focuses on introducing the basics of Go programming language. Whether you have prior experience in languages
like C, Python, or Java, this module will help you grasp the fundamental elements of Go and develop your programming
skills further.

| Assignment | File Name                                                                                                     | Description                                                                                                                                    |
|------------|---------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| 1          | [hello.go](./1-getting-started/assignment-1/hello.go)                                                         | A simple program that prints "Hello, world!" to the console.                                                                                   |
| 2          | [trunc.go](./1-getting-started/assignment-2/trunc.go)                                                         | A program that reads a floating-point number from the user and prints its truncated integer value.                                             |
| 3          | [findian.go](./1-getting-started/assignment-3/findian.go)                                                     | A program that checks if a user-provided string contains the letters 'i', 'a', and 'n' in any order.                                           |
| 4          | [slice.go](./1-getting-started/assignment-4/slice.go)                                                         | A program that reads integers from the user until 'X' is entered, sorts them in ascending order, and prints the sorted slice after each input. |
| 5          | [makejson.go](./1-getting-started/assignment-5/makejson.go)                                                   | A program that prompts the user to enter a name and an address, creates a JSON object with the provided data, and prints it to the console.    |
| 6          | [read.go](./1-getting-started/assignment-6/read.go) and [data.txt](./1-getting-started/assignment-6/data.txt) | A program that reads names from a text file (specified by the user) and prints them along with their first and last names.                     |

## Functions, Methods, and Interfaces in Go

The module aims to deepen the understanding of the Go programming language by exploring functions, methods, and
interfaces. The topics covered include the implementation of functions, function types, object-orientation in Go,
methods, and class instantiation.

<details>
  <summary>Assignment 1 - Bubble Sort </summary>

# Integer Bubble Sort

The Integer Bubble Sort program is a command-line application that allows users to enter up to 10 integers separated by
spaces. It then sorts the integers using the bubble sort algorithm and displays the sorted array.

## Usage

1. Run the program.
2. Enter up to 10 integers separated by spaces when prompted.
3. The program will validate the input and sort the integers using the bubble sort algorithm.
4. The sorted array will be displayed.

## Algorithm

The program uses the bubble sort algorithm to sort the integers. The bubble sort algorithm works by repeatedly stepping
through the list, comparing adjacent elements, and swapping them if they are in the wrong order. This process is
repeated until the list is sorted.

## Example

```
Enter up to 10 integers (separated by spaces):
> 9 5 2 7 1

Array after sorting: [1 2 5 7 9]
```

## Limitations

- The program only accepts up to 10 integers. If more than 10 numbers are entered, the program will truncate the input
  to contain only the first 10 numbers.
- The program assumes valid input, where each input is a valid integer separated by spaces. It does not handle
  non-integer inputs or inputs with incorrect format.

</details>

<details>
  <summary>Assignment 2 - Displacement Calculator </summary>

# Displacement Calculator

The Displacement Calculator is a command-line application that allows users to calculate displacement based on time,
acceleration, initial velocity, and initial displacement. The program prompts the user to enter values for acceleration,
initial velocity, and initial displacement, and then computes the displacement after a specified time.

## Usage

1. Run the program.
2. Enter the values for acceleration, initial velocity, and initial displacement when prompted.
3. Enter the desired time for which you want to calculate the displacement.
4. The program will calculate the displacement using the provided values and display the result.

## Formula

The displacement calculation is based on the following formula:

```
displacement = 0.5 * acceleration * time^2 + initialVelocity * time + initialDisplacement
```

## Example

```
Enter acceleration:
2.5
Enter initial velocity:
1.0
Enter initial displacement:
-3.0
Enter time:
4.0

Displacement after 4 seconds: 37.0
```

## Limitations

- The program assumes valid numeric input for acceleration, initial velocity, initial displacement, and time. It does
  not handle non-numeric inputs or inputs with incorrect format.

</details>

<details>
  <summary>Assignment 3 - OOP in Go</summary>

A program that allows the user to interactively query information about different animals. It demonstrates the usage of
maps, structs, and methods in Go.

The program defines a set of predefined animals with their attributes (food, locomotion, and noise). It prompts the user
to enter an animal name and the information they want to know (eat, move, or speak). Then, it displays the corresponding
information about the animal.

The table below showcases a variety of animals along with their respective characteristics.

| Animal | Food Eaten | Locomotion Method | Spoken Sound |
|--------|------------|-------------------|--------------|
| cow    | grass      | walk              | moo          |
| bird   | worms      | fly               | peep         |
| snake  | mice       | slither           | hsss         |

Once the program is running, you can start entering commands. Each command should consist of an animal name and the
desired information, separated by a space.

Here's an example of how you can interact with the program:

   ```
   > cow eat
   Grass
   > bird move
   Fly
   ```

</details>

<details>
  <summary>Assignment 4 - Polymorphism in Go</summary>

The fourth assignment extends the functionality of the previous assignment by allowing the user to dynamically create
new animals and add them to the existing set of animals.

The table below showcases a variety of animals along with their respective characteristics.

| Animal | Food Eaten | Locomotion Method | Spoken Sound |
|--------|------------|-------------------|--------------|
| cow    | grass      | walk              | moo          |
| bird   | worms      | fly               | peep         |
| snake  | mice       | slither           | hsss         |

The program prompts the user to enter a command ("newanimal" or "query"), an animal name, and additional details based
on the command.

* For the "`newanimal`" command, it creates a new animal of the specified type and adds it to the animals map.
* For the "`query`" command, it retrieves the requested information about the specified animal.

To interact with the program, use the following format:

  ```
> newanimal <animal_name> <details>
> query <animal_name> <details>
  ```

For the `newanimal` command, the `details` field represents the type of the animal to create (e.g., "cow", "bird", "
snake").

For the `query` command, the `details` field represents the information requested about the animal (e.g., "eat", "
move", "speak").

Here's an example of how you can interact with the program:

  ```
> newanimal cow moomoo
> Created it!
> newanimal bird tweetie
> Created it!
> query cow move
> walk
> query bird speak
> peep
  ```

</details>

## Concurrency in Go

This module provides an introduction to concurrent programming in Go, focusing on the roles of channels and goroutines
in implementing concurrency. The topics covered include writing goroutines and implementing channels for communication
between goroutines.

<details>
  <summary>Assignment 1 - Race conditions</summary>

This Go program showcases a race condition scenario when two goroutines are executed concurrently. It provides a simple
demonstration of how a race condition can occur and affect the outcome of the program.

The race condition in this program occurs due to the concurrent access and modification of the shared variable `counter`
by the `incrementCounter` and `printCounter` goroutines. Both goroutines read and write to the counter variable without
any synchronization mechanism.

During execution, the `incrementCounter` goroutine reads the current value of `counter`, increments it, and writes the
updated value back. At the same time, the `printCounter` goroutine reads the value of `counter` and prints it.

Since there is no synchronization between the goroutines, the interleaving of their operations can lead to unexpected
results. The race condition manifests when both goroutines read the counter value simultaneously, resulting in
inconsistencies in the printed output.

</details>

<details>
  <summary>Assignment 2 - Parallel merge sort</summary>

This program implements parallel merge sort to efficiently sort a series of integers. It takes input from the user,
which should be a space-separated series of integers. The program divides the input into four subarrays and uses
goroutines to sort each subarray concurrently. Finally, it merges the sorted subarrays to obtain the sorted array.

By dividing the input into four subarrays and sorting them concurrently using goroutines, the program takes advantage of
parallelism to improve performance. However, the performance gain depends on the size of the input and the hardware
resources available.

Note that the program assumes the input contains a valid series of integers separated by spaces. It performs minimal
error checking and may produce unexpected results if the input is not properly formatted.

</details>

### Workspace Configuration

The repository also includes a Visual Studio Code workspace configuration file (`programming-with-go.code-workspace`).
You can open the entire project in Visual Studio Code using this workspace file.

---

## Running the Project

To run the assignments in the "Getting Started with Go" module, follow these steps:

1. Install Go on your machine. You can download it from the official Go
   website: [https://golang.org/dl/](https://golang.org/dl/).

2. Clone or download this repository to your local machine.

3. Open a terminal or command prompt and navigate to the root directory of the cloned repository.

4. To run a specific assignment, use the `go run` command followed by the path to the Go file. For example, to
   run `hello.go`, use the following command:

   ```bash
   go run ./1-getting-started/assignment-1/hello.go
   ```

Replace the file path with the path to the desired assignment.

Note: Some assignments may require additional files (e.g., `data.txt`). Make sure to provide the necessary input
files when prompted.

5. Follow the instructions or input the required values when prompted by the program.

---

Feel free to explore and run these assignments to learn Go programming!
