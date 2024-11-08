<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";

    let email = "";
    let password = "";
    let message = "";

    const apiUrl = "http://localhost:8080/api/users/login";

    const handleSubmit = async () => {
        try {
            const response = await fetch(apiUrl, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    email,
                    password,
                }),
                credentials: "include",
            });

            if (response.ok) {
                message = "Login successful!";
                const data = await response.json();
                goto("/profile");

                await checkLogin();
            } else {
                message = `Failed: ${response.statusText} (${response.status})`;
            }
        } catch (error) {
            message = "Error submitting the form: " + error.message;
        }
    };

    const checkLogin = async () => {
        try {
            const response = await fetch(
                "http://localhost:8080/api/users/profile",
                {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    credentials: "include",
                },
            );

            if (response.ok) {
                const data = await response.json();
                console.log(data);
                sessionStorage.setItem("profileData", JSON.stringify(data));

                message = "Login successful!";
                goto("/profile");
            } else {
                console.error("Profile check failed:", response.statusText);
            }
        } catch (error) {
            console.error("Error during profile check:", error);
        }
    };

    onMount(() => {
        document.body.classList.add("dark");
        checkLogin();
    });
</script>

<div class="wrapper">
    <div class="centered-div">
        <h1 class="logo">YourMusic 🎶</h1>
        <div class="form-class">
            <form on:submit|preventDefault={handleSubmit}>
                <h2>Sign in to YourMusic</h2>
                <p>Email</p>
                <input
                    type="email"
                    bind:value={email}
                    placeholder="Email"
                    required
                />
                <p>Password</p>
                <input
                    type="password"
                    bind:value={password}
                    placeholder="Password"
                    required
                />
                <button class="rounded-button" type="submit">Login</button>

                {#if message}
                    <p>{message}</p>
                {/if}
            </form>
        </div>
        <div>
            <a href="/register">Forgot your password?</a>
        </div>
        <div>
            Don't have an account? <a href="/register">Register here!</a>
        </div>
    </div>
</div>
