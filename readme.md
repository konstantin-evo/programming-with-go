## Programming with Google Go Specialization

This repository contains the code for the "Programming with Google Go Specialization" on Coursera. The specialization
consists of several modules, each focusing on different aspects of programming in Go.

<p align="center">
  <img src="src/go-logo.png" alt="Go Logo" width="30%">
</p>

### Modules

1. **Getting Started with Go**

## Assignment Descriptions

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

2. **Functions, Methods, and Interfaces in Go**

(Upcoming)

3. **Concurrency in Go**

(Upcoming)

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
