<script>
    import { onMount } from "svelte";
    let email = "";
    let id = "";
    const handleToken = async () => {
        const token = localStorage.getItem("token");

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
            });
    };

    onMount(() => {
        handleToken();
    });
</script>

<main>
    <p>{email}</p>
    <p>{id}</p>
</main>

<style>
    main {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        width: 100%;
        height: 100%;
        background-color: rgb(0, 0, 0);
    }
    p {
        color: white;
    }
</style>
