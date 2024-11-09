<script>
</script>

<main>
    <div class="signup">
        <h1 class="logo">yourMusic 🎶</h1>
        <p class="logo">Sign Up</p>
        <form action="">
            <label for="username">Username:</label>
            <input type="username" id="username" />
            <label for="email">Email:</label>
            <input type="email" id="email" />
            <label for="password">Password:</label>
            <input type="password" id="password" />
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
