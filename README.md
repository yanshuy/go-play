# Custom TCP Chat Server

This is a simple TCP-based chat server implemented in Go. It allows multiple clients to connect, set nicknames, join chat rooms, send messages, and list available rooms.

## Features

-   **Nicknames**: Clients can set their own nicknames using the `/nick` command.
-   **Chat Rooms**: Clients can create or join chat rooms using the `/join` command.
-   **Broadcast Messaging**: Messages sent in a room are broadcast to all members of the room.
-   **Room Listing**: Clients can list all available rooms using the `/rooms` command.
-   **Graceful Exit**: Clients can leave the chat server using the `/quit` command.

## Commands

| Command        | Description                              |
| -------------- | ---------------------------------------- |
| `/nick <name>` | Set your nickname.                       |
| `/join <room>` | Join or create a chat room.              |
| `/rooms`       | List all available chat rooms.           |
| `/msg <text>`  | Send a message to the current chat room. |
| `/quit`        | Disconnect from the server.              |

## How to Run

1. Clone the repository:

    ```bash
    git clone <repository-url>
    cd custom_tcp
    ```

2. Build and run the server:

    ```bash
    go build -o myapp
    ./myapp
    ```

3. Connect to the server using a TCP client (e.g., `telnet` or `nc`):

    ```bash
    telnet localhost 4000
    ```

4. Use the commands listed above to interact with the server.

## Project Structure

-   `main.go`: Entry point of the application.
-   `server.go`: Handles server operations and client connections.
-   `client.go`: Manages client interactions and commands.
-   `room.go`: Defines chat room behavior and member management.
-   `commands.go`: Defines command types and IDs.

## Example Usage

1. Start the server:

    ```bash
    ./myapp
    ```

2. Connect multiple clients using `telnet` or `nc`:

    ```bash
    telnet localhost 4000
    ```

3. Set a nickname:

    ```
    /nick yanshuy
    ```

4. Join or create a room:

    ```
    /join general
    ```

5. Send a message:

    ```
    /msg Hello, everyone!
    ```

6. List available rooms:

    ```
    /rooms
    ```

7. Quit the server:
    ```
    /quit
    ```

## License

This project is licensed under the MIT License.
