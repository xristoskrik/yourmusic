<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    let email = "";
    let password = "";

    const handleToken = async (event) => {
        const token = localStorage.getItem("token");
        if (token === "undefined" || token === null) {
            console.log(token);
            return;
        }

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
                console.log(response.json());
                goto("/profile");
            })
            .then((data) => {
                console.log("Profile:", data);
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        const payload = {
            email,
            password,
        };

        try {
            const response = await fetch(
                "http://localhost:8080/api/users/login",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(payload),
                },
            );

            if (response.ok) {
                const result = await response.json();
                console.log(result.token);
                localStorage.setItem("token", result.token);

                goto("/profile");
                console.log(result);
            } else {
                const error = await response.json();
                alert("Error: " + error.message);
            }
        } catch (error) {
            console.error("Error:", error);
            alert("Can't log in.");
        }
    };
    onMount(() => {
        handleToken();
    });
</script>

<main>
    <div class="login">
        <h1 class="logo">yourMusic 🎶</h1>
        <p class="logo">Log In</p>
        <form on:submit|preventDefault={handleSubmit}>
            <label for="username">Username or email:</label>
            <input type="username" id="username" bind:value={email} />
            <label for="password">Password:</label>
            <input type="password" id="password" bind:value={password} />
            <button>Log In!</button>
        </form>
    </div>
</main>

<style>
    main {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: #121212;
    }
    .login {
        width: 40%;
        height: 70%;
        transition: width 0.5s ease-in-out;
        color: white;
        background-color: #1c1c1e;
        border: 8px solid #0d0d0d;
        border-radius: 3vh;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    p {
        font-size: 1.5rem;
    }
    h1 {
        margin-top: 2rem;
        font-size: 4vh;
        font-family: roboto;
        letter-spacing: 0.3vh;
    }

    form {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    label {
        margin-bottom: 1vh;
    }
    button {
        margin-top: 4vh;
        padding: 1vh;
        width: 100%;
        background-color: green;
        border: 5px solid green;
        border-radius: 3vh;
    }
    input {
        width: 100%;
        padding: 1vh;
        margin-bottom: 1vh;
    }
    @media (max-width: 600px) {
        .login {
            width: 70%;
        }
    }
</style>
