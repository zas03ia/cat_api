document.getElementById("fav").addEventListener("click", () => {
    const imageID = "{{.ID}}"; // Dynamically passed from the template

    // Send a POST request to the backend to favorite an image
    fetch('/favourite', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ image_id: imageID }),
    })
    .then(response => {
        if (response.ok) {
            window.location.href = '/'; // Redirect to the home route on success
        } else {
            console.error("Failed to favourite the image.");
        }
    })
    .catch(error => console.error('Error:', error));
});
