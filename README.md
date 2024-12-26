# Cat Image Viewer Project

This project is a web-based application built using **Beego**, a Go web framework, to display images of cats, manage favourite images, and allow users to vote on images. The app fetches data from The Cat API and provides functionality like showing random cat images, categorizing them by breeds, allowing users to add images to their favourites, and vote on them.


## Features
- Display random cat images.
- Fetch and display cat breeds.
- Fetch breed-specific images.
- Add images to favourites.
- Vote on images (Upvote/Downvote).
- Handle concurrent fetching of data from multiple sources (images, breeds, favourites).

## Tech Stack
- **Backend Framework**: Beego (Go)
- **Frontend**: HTML, CSS, JavaScript (Bootstrap, Font Awesome)
- **Database**: None (All data is fetched from The Cat API)
- **API**: The Cat API (https://thecatapi.com/)

## Installation

### Prerequisites
1. **Go**: Install Go (version 1.18 or later). Follow the installation instructions here: [Go Installation](https://golang.org/doc/install).
2. **Beego**: Install Beego framework by running the following command:
   ```bash
   go install github.com/beego/beego/v2@latest
   ```

3. **API Key**: Sign up at [The Cat API](https://thecatapi.com/) to get an API key. You will need to add your API key to the `app.conf` file.

### Clone the Repository
Clone this repository to your local machine:
```bash
git clone [https://github.com/your-username/cat-image-viewer.git](https://github.com/zas03ia/cat_api.git)
```

### Install Dependencies
In the project directory, run the following command to install dependencies:
```bash
go mod tidy
```

### Configuration
1. Create an `app.conf` file in the root of the project and add your API key:
   ```ini
   api_key = your_api_key_here
   get_images_api_url = https://api.thecatapi.com/v1/images/search
   ```

## Running the Application

### Run the Application
To start the application, run the following command in the root project directory:
```bash
bee run
```

This will start the Beego server at `http://localhost:8080/`.

### Accessing the Application
- Open your browser and go to `http://localhost:8080/` to view the cat images, vote, and manage favourites.

## Folder Structure
The project has the following folder structure:
```
├── controllers/            # Controllers to handle API requests and business logic
├── models/                # Data models for API responses
├── routers/               # Application routes
│   └── routers.go         # Setup all routes for the application
├── static/                # Static files (e.g., CSS, JS)
│   ├── css/
│   ├── js/
├── views/                 # Templates (HTML)
│   └── index.tpl          # Main template for displaying cat images
├── app.conf               # Application configuration file (API Key, URLs)
├── go.mod                 # Go module file
└── README.md              # Project documentation
```

## API Endpoints
### 1. `/`
- **Method**: `GET`
- **Description**: Displays random cat images, breed information, and favourites.

### 2. `/vote`
- **Method**: `POST`
- **Description**: Allows users to vote on images (Upvote/Downvote).
- **Parameters**: `image_id` (ID of the image), `vote` (1 for upvote, 0 for downvote)

### 3. `/favourite`
- **Method**: `POST`
- **Description**: Allows users to add an image to their favourites.
- **Parameters**: `image_id` (ID of the image)

### 4. `/favourites`
- **Method**: `GET`
- **Description**: Fetches all favourite images.

## Testing

To test the application, you can use unit testing in Go with Beego's built-in test suite.

### Running Tests
To run the tests, use the following command:
```bash
go test ./...
```
