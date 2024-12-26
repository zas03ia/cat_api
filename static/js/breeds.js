
const breedsDropdown = document.getElementById("breeds");



function setSelectedBreed() {
    var breedId = document.getElementById("breeds").value;
    if (breedId) {
        // Ensure the breedId is a string and set it as the selected option
        breedsDropdown.value = String(breedId);
        breed_info = document.querySelector(".breed_info");
        const selectedBreed = breeds.find(breed => breed.id === breedId);

        if (selectedBreed) {
            document.querySelector(".breed_info").innerHTML = `
                <p class="title">${selectedBreed.name} <span>(${selectedBreed.origin})</span> <em>${selectedBreed.id}</em></p>
                <p class="description">${selectedBreed.description}</p>
                <a class="wiki" href="${selectedBreed.wikipedia_url}" target="_blank">WIKIPEDIA</a>`;
        } else {
            document.querySelector(".breed_info").innerHTML = "<p>No description available.</p>";
        }
    }
}

// Slideshow functionality
const slideshow = document.getElementById("slideshow");
let slideIndex = 0;

let slideTimeout; // Global variable to store the timeout ID

function showSlides() {
    const slides = document.querySelectorAll(".slide");
    const dots = document.querySelectorAll(".dot");

    slides.forEach((slide, index) => {
        slide.style.display = index === slideIndex ? "block" : "none";
    });

    dots.forEach((dot, index) => {
        dot.classList.remove("active");
        if (index === slideIndex) {
            dot.classList.add("active");
        }
    });

    slideIndex = (slideIndex + 1) % slides.length; // Loop back to the start

    clearTimeout(slideTimeout); // Clear any existing timeout
    slideTimeout = setTimeout(showSlides, 3000); // Start a new timeout
}
// Create dots dynamically based on the number of slides
function createDots() {
    const slides = document.querySelectorAll(".slide");
    const dotsContainer = document.getElementById("dots");
    dotsContainer.innerHTML = "";

    slides.forEach((slide, index) => {
        const dot = document.createElement("span");
        dot.classList.add("dot");
        dot.addEventListener("click", () => {
            slideIndex = index;
            showSlides(); // Show the selected slide
        });
        dotsContainer.appendChild(dot);
    });
    slideIndex = 0;
}

function updatedata() {

    const breedID = document.getElementById("breeds").value;

    const slideshowContainer = document.getElementById("slideshow");

    // Clear the slideshow container
    slideshowContainer.innerHTML = "";
    if (breedID) {

        const selectedBreedImages = breed_images[breedID]; // Fetch images for the selected breed

        if (selectedBreedImages && selectedBreedImages.length > 0) {

            selectedBreedImages.forEach(function (image, index) {

                const slide = document.createElement("div");
                slide.className = "slide";

                const img = document.createElement("img");
                img.src = image.url;
                img.alt = "Cat";

                slide.appendChild(img);
                slideshowContainer.appendChild(slide);
            });
            setSelectedBreed();
            createDots(); // Update dots for the new slides
            showSlides(); // Restart slideshow
            
        } else {
            alert("No images found for the selected breed."); // Step 12: No images case
        }
    } else {
        alert("No Breed ID selected."); // Step 13: No selection case
    }
}


updatedata()