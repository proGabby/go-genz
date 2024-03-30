# GENZ: Microblogging Service

## Overview
GENZ is a microblogging service designed to provide users with a platform for sharing posts, reacting to posts, and engaging in real-time chat conversations. Built with Go, PostgreSQL, WebSocket, and AWS S3, GENZ offers a modern and efficient solution for users to interact with each other in a dynamic and engaging environment.

## Key Features
1. **Feed Creation:** Users can create posts (feeds) with captions and attach images to express themselves and share content with others.
2. **Feed Interaction:** Users can react to feeds with various emoticons or engage in discussions through comments.
3. **Real-Time Chat:** GENZ includes real-time chat functionality, allowing users to exchange messages with each other instantly.
4. **User Authentication:** Secure user authentication ensures that only registered users can access and interact with the platform.
5. **File Storage:** GENZ utilizes AWS S3 for storing user-uploaded images, ensuring reliable and scalable file storage.
6. **Clean Architecture:** The application is built following clean architecture principles, ensuring modularity, maintainability, and scalability.

## Technologies Used
- **Go (Golang):** Backend development is powered by Go, offering high performance, concurrency, and efficiency.
- **PostgreSQL:** As the database of choice, PostgreSQL provides robust data storage and retrieval capabilities.
- **WebSocket:** WebSocket technology enables real-time bidirectional communication between the server and clients for chat functionality.
- **AWS S3:** Amazon S3 is used for storing user-uploaded images securely and reliably.
- **Clean Architecture:** The application architecture is organized following clean architecture principles, separating concerns into layers for better maintainability and testability.

## Getting Started
1. **Prerequisites:** Ensure you have Go, PostgreSQL, and AWS S3 credentials set up.
2. **Installation:** Clone the GENZ repository and install dependencies.
3. **Configuration:** Set up environment variables and configuration files for database connection and AWS S3 credentials.
4. **Run:** Start the GENZ server and ensure all services are up and running.
5. **Usage:** Access the GENZ web interface or API endpoints to create feeds, react to feeds, and engage in chat conversations.
