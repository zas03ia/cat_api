<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Image Viewer</title>
    <link rel="stylesheet" href="/static/css/base.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap-icons/font/bootstrap-icons.css" rel="stylesheet">
</head>
<body>

    <script>
        // Inject dynamic data from Go into JavaScript
        const imageID = "{{.imageID}}"; // Dynamically set the imageID from Go
        const breeds = JSON.parse("{{.Breeds_json}}");
        const breed_images = JSON.parse("{{.breedImages}}");
    </script>
    <div class="first">
        <h1>The Cat API <hr>
        <strong>Cats as a service.</strong></h1>
        
    </div>

    <div class="card">
        <!-- Header Navigation -->
        <div class="header">
            <div id="vote-nav" class="nav active" onclick="navigateTo('voting')">
                <span class="header-icon">&#8679;&#8681;</span><br>
                <span><em> Voting</em></span>
            </div>
            <div id="breed-nav" class="nav" onclick="navigateTo('breeds')">
                <span class="header-icon">&#9740;</span><br>
                <span><em>Breeds</em></span>
            </div>
            <div id="fav-nav" class="nav" onclick="navigateTo('favs')">
                <span class="header-icon">&#9825;</span><br>
                <span> <em>Favs</em></span>
            </div>
        </div>

        <!------------------------------- Voting -------------------------------->
        <div class="changeable1">
            <link rel="stylesheet" href="/static/css/voting.css">
            <div class="image-container">
                {{if .imageURL}}
                    <img src="{{.imageURL}}" alt="Cat Image" />
                {{else}}
                    <p>No image found</p>
                {{end}}

            </div>

            <!-- Footer Like, Dislike and Favorite Buttons -->
            <div class="footer">
                <!-- Favorite Button -->
                <form action="/favourite" method="POST">
                    <input type="hidden" name="image_id" value="{{.imageID}}">
                    <button type="submit" title="Favorite" class="fav-icon">&#9825;</button>
                </form>

                <!-- Like and Dislike Buttons -->
                <div class="like-dislike">
                    <!-- Like Button -->
                    <form action="/vote" method="POST">
                        <input type="hidden" name="image_id" value="{{.imageID}}">
                        <input type="hidden" name="vote" value="1">
                        <button type="submit" title="Like"><i class="fa-regular fa-thumbs-up"></i></button>
                    </form>

                    <!-- Dislike Button -->
                    <form action="/vote" method="POST">
                        <input type="hidden" name="image_id" value="{{.imageID}}">
                        <input type="hidden" name="vote" value="-1">
                        <button type="submit" title="Dislike"><i class="fa-regular fa-thumbs-down"></i></button>
                    </form>
                </div>
            </div>

        </div>
        <script src="/static/js/make_favorite.js" defer></script>
        <script src="/static/js/vote.js" defer></script>
        <!--------------------------------Breeds---------------------------------------->
        <div class="changeable2">
            <link rel="stylesheet" href="/static/css/breed.css">
            <div class="dropdown-container">
                <select class="options" id="breeds" onchange="updatedata()">
                    {{range $index, $breed := .breeds}}
                    <option value="{{$breed.ID}}" {{if eq $index 0}}selected{{end}}>{{$breed.Name}}</option>
                    {{end}}
                </select>
            </div>
            
            <div class="slideshow-container" id="slideshow-container">
                <div id="slideshow">
                  
                </div>
                <div class="dots-container" id="dots"></div> <!-- Dots container added here -->
            </div>
            <div class="breed_info"></div>

            <script src="/static/js/breeds.js" defer></script>
        </div>


        <!------------------------------------Favourites----------------------------------->
       <div class="changeable3">
            <link rel="stylesheet" href="/static/css/favorites.css">

            <div class="view-toggle">
                <button id="gridViewButton" onclick="setGridView()"><i class="fa-solid fa-grip"></i></button>
                <button id="linearViewButton" onclick="setLinearView()"><i class="fa-solid fa-grip-lines"></i></button>
            </div>

            <div class="favourites-container">
                {{if .favourites}}
                    {{range $fav := .favourites}}
                        <div class="favourite">
                            <img src="{{$fav.Image.URL}}" alt="Cat Image" />
                        </div>
                    {{end}}
                {{else}}
                    <p>No favourites found.</p>
                {{end}}
            </div>
            <script src="/static/js/show_favourites.js" defer></script> 
        </div>


    </div>

    <script src="/static/js/index.js" defer></script>
</body>
</html>