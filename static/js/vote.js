function reloadImage() {
    location.reload();
}


// Function to handle voting (Like/Dislike)
function vote(value) {
    const imageID = "{{.ID}}";

    const data = {
        image_id: imageID,
        sub_id: "unique_user_id",
        value: value
    };

    // Send a POST request to the backend to vote
    fetch('/vote', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {
        // If successful, redirect to the home route
        window.location.href = '/';
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

