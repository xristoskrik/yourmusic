<script>
    let username = "";
    let email = "";
    let password = "";

    const handleSubmit = async (event) => {
        event.preventDefault();

        const payload = {
            username,
            email,
            password,
        };

        try {
            const response = await fetch("http://localhost:8080/api/users", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(payload),
            });

            if (response.ok) {
                const result = await response.json();
                alert("Sign-up successful!");
                console.log(result);
            } else {
                const error = await response.json();
                alert("Error: " + error.message);
            }
        } catch (error) {
            console.error("Error:", error);
            alert("An error occurred while signing up.");
        }
    };
</script>

<main>
    <div class="signup">
        <h1 class="logo">yourMusic 🎶</h1>
        <p class="logo">Sign Up</p>
        <form on:submit|preventDefault={handleSubmit}>
            <label for="username">Username:</label>
            <input type="username" id="username" bind:value={username} />
            <label for="email">Email:</label>
            <input type="email" id="email" bind:value={email} />
            <label for="password">Password:</label>
            <input type="password" id="password" bind:value={password} />
            <button>Sign Up!</button>
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
    .signup {
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
        .signup {
            width: 70%;
        }
    }
</style>
