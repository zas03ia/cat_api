const vote_nav = document.getElementById("vote-nav");
const breed_nav = document.getElementById("breed-nav");
const fav_nav = document.getElementById("fav-nav");

const c1 = document.querySelector(".changeable1");
const c2 = document.querySelector(".changeable2");
const c3 = document.querySelector(".changeable3");

function navigateTo(section) {
    vote_nav.classList.remove("active");
    breed_nav.classList.remove("active");
    fav_nav.classList.remove("active");
    c1.style.display = "none";
    c2.style.display = "none";
    c3.style.display = "none";
    
    // Determine the URL and class based on the selected section
    switch(section) {
        
        case 'voting': // (class=changeable1)
            vote_nav.classList.add("active");
            c1.style.display = "block";
            break;

        case 'breeds': // (class=changeable2)
            breed_nav.classList.add("active");
            c2.style.display = "block";
            break;

        case 'favs': // (class=changeable3)
            fav_nav.classList.add("active");
            c3.style.display = "block";
            break;

        default:
            return;
    }
}
