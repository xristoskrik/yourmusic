<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    let email = "";
    let id = "";
    const handleAlbum = async () => {
        console.log("go to album");
    };

    const handleToken = async () => {
        const token = localStorage.getItem("token");
        if (token !== null) {
            fetch("http://localhost:8080/api/users/profile", {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
                credentials: "include",
            })
                .then((response) => {
                    if (!response.ok) {
                        throw new Error("Failed to fetch profile");
                    }
                    return response.json();
                })
                .then((data) => {
                    console.log("Profile:", data);
                    email = data.Email;
                    id = data.ID;
                })
                .catch((error) => {
                    console.error("Error:", error);
                    goto("/login");
                });
        } else {
            goto("/login");
        }
    };

    onMount(() => {
        handleToken();
    });
    const artists = [
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
        { id: "1", name: "Artist", img: "./image.png" },
    ];
</script>

<main>
    <!--
    <p>{email}</p>
    <p>{id}</p>-->
    <div class="artists">
        <p class="artists-header">Artists</p>
        {#each artists as artist}
            <a href="/albums/{artist.id}" class="artist">
                <img src={artist.img} alt="artist" />
                <p class="song-details">{artist.name}</p>
            </a>
        {/each}
    </div>
</main>

<style>
    main {
        display: flex;
        justify-content: center;
        width: 100%;
        height: 100%;
        background-color: rgb(0, 0, 0);
    }
    .artists {
        width: 90%;
        height: 100%;
        display: flex;
        flex-wrap: wrap;
        overflow-y: auto;
        justify-content: center;
        align-content: flex-start;
        background-color: blue;
        padding: 1rem;
        box-sizing: border-box;
    }

    .artist {
        display: flex;
        align-items: center;
        flex-direction: column;
        margin: 0.5rem;
        width: 20%;
        margin-top: 2em;
    }
    .song-details {
        font-size: small;
    }
    img {
        width: 70%;
        margin-right: 0.5rem;
    }
    .artists-header {
        width: 100%;
        text-align: center;
        color: white;
        margin-bottom: 1rem;
        font-size: 1.5rem;
    }
    p {
        color: white;
        margin: 0;
    }
    .artist:hover {
        background-color: red;
    }

    @media (max-width: 600px) {
        .artist {
            width: 100%;
            margin: 0.5rem 0;
            overflow-y: auto;
            overflow-x: hidden;
        }
    }
</style>
