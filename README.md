# simple-socket-server 🌐

A minimal, proof-of-concept **TCP socket server** implementation written in the **Go** programming language. This repository provides a basic foundation for handling incoming client connections and simple message exchange over a network socket.

---

## Prerequisites

To run this server, you must have:

* **Go** (Golang) installed on your system.
* **`git`** for cloning the repository.
* **`nc`** (netcat) installed for testing the connection.

---

## How to Run the Server

Follow these steps to clone the repository and start the server:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/ProArash/simple-socket-server.git
    cd simple-socket-server
    ```

2.  **Start the server:**
    Run the main Go file. (The server is assumed to listen on **`localhost:4666`**.)
    ```bash
    go run main.go
    ```
    The server should now be running and waiting for client connections.

---

## How to Test with `nc` (Netcat)

Use the `nc` command in a separate terminal window to connect to the running server.

1.  **Connect to the server:**
    ```bash
    nc localhost 8080
    ```

2.  **Send a message:**
    Once connected, you can type a message and press **Enter**. The server should typically process the message and may send a response back (depending on its implementation).

    **Example Interaction:**
    ```
    nc localhost 8080
    Hello, server!
    [Server Response Here]
    ```

    *Note: To close the client connection, press **Ctrl+C** or **Ctrl+D**.*