# 🐜 Lem-in: Ant Colony Simulation

Lem-in is a simulation project where a group of ants must travel from a starting room to an ending room through a network of tunnels. The goal is to optimize their movements to minimize the number of turns required.

---

## 📌 Features
✅ Reading and analyzing input data (number of ants, rooms, links).  
✅ Finding possible paths between the start and end rooms.  
✅ Intelligent distribution of ants to optimize their journey.  
✅ Displaying the movement of ants at each turn.  

---

## 📁 Project Structure
```
lem-in/
│── cmd/
│   └── main.go         # Program entry point
│── bfs/
│   └── bfs.go          # Pathfinding algorithm (BFS)
│── parsing/
│   └── parsing.go      # Reading and analyzing input file
│── printage/
│   └── printage.go     # Displaying ant movements
│── utils/
│   └── utils.go        # Utility functions and data structures
│── README.md           # Documentation
```

---

## 📥 Installation and Execution
### 1️⃣ Prerequisites
- **Go** installed on your machine (check with `go version`).

### 2️⃣ Clone the Project
```sh
git clone https://github.com/RedaAz07/lem-in.git
cd lem-in
```

### 3️⃣ Compile and Run
```sh
go run . test0.txt
```
Replace `test0.txt` with the file containing the ant farm description.

---

## 📜 Input File Format
The input file must follow this format:
```
3               # Number of ants
##start         # Start of room definitions
A 1 2          # Starting room
B 3 4
C 5 6
##end           # Start of the ending room definition
D 7 8          # End room
A-B            # Links between rooms
A-C
B-D
C-D
```

---

## 🔎 How It Works
1. **Reading the input file** (`parsing.go`):
   - Extracting the number of ants, rooms, and tunnels.
2. **Finding optimal paths** (`graph.go`):
   - BFS algorithm to find the best paths.
3. **Simulating ant movements** (`printage.go`):
   - Distributing ants across paths and displaying their movements.

---

## 📖 Execution Example
**Input (test0.txt):**
```
3
##start
A 1 2
B 3 4
C 5 6
##end
D 7 8
A-B
A-C
B-D
C-D
```

**Expected Output:**
```
L1-B L2-C
L1-D L2-D L3-B
L3-D
```

---

## 🤝 Contribution
Contributions are welcome! To propose improvements:
1. **Fork** the project.
2. Create a branch (`git checkout -b feature-new-feature`).
3. Make your changes and commit (`git commit -m "Added a new feature"`).
4. Push your changes (`git push origin feature-new-feature`).
5. Create a Pull Request.

---

## 🛠 Technologies Used
- **Language:** Go
- **Main Algorithm:** BFS (Breadth-First Search)
- **File Management:** `os` and `strings`

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

