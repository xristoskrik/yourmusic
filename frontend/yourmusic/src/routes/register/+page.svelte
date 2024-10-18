<script>
    let email = ""; // Replaced username with email
    let password = "";
    let apiUrl = "http://localhost:8080/api/users"; // The URL the user will input or you can hardcode this
    let message = "";

    const handleSubmit = async () => {
        try {
            const response = await fetch(apiUrl, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    email, // Send email instead of username
                    password,
                }),
            });

            if (response.status === 201) {
                const data = await response.json();
                message = `Success: ${JSON.stringify(data)}`; // Show the JSON response
            } else {
                message = `Failed: ${response.statusText} (${response.status})`;
            }
        } catch (error) {
            message = "Error submitting the form: " + error.message;
        }
    };
</script>

<form on:submit|preventDefault={handleSubmit}>
    <p>{message}</p>
    <input type="email" bind:value={email} placeholder="Email" />
    <!-- Changed to email -->
    <input type="password" bind:value={password} placeholder="Password" />
    <button type="submit">Register</button>
</form>
