const gridViewButton = document.getElementById("gridViewButton");
const linearViewButton = document.getElementById("linearViewButton");
const favouritesContainer = document.querySelector(".favourites-container"); // Assuming this is the container class

function setGridView() {
    favouritesContainer.classList.remove("linear-view");
    favouritesContainer.classList.add("grid-view");

    // Update button states
    gridViewButton.classList.add("active");
    linearViewButton.classList.remove("active");
}

function setLinearView() {
    favouritesContainer.classList.remove("grid-view");
    favouritesContainer.classList.add("linear-view");

    // Update button states
    linearViewButton.classList.add("active");
    gridViewButton.classList.remove("active");
}


// Default to grid view on page load
setGridView();
