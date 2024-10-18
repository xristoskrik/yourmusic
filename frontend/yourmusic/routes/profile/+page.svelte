<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    let profileData = {};
    let email = "";
    let id = "";
    const logout = async () => {
        try {
            const response = await fetch(
                "http://localhost:8080/api/users/logout",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    credentials: "include",
                },
            );
            console.log(response);
            if (response.ok) {
                goto("/");
            } else {
                console.error(
                    `Failed: ${response.statusText} (${response.status})`,
                );
            }
        } catch (error) {
            console.error("Error submitting the form: " + error.message);
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
                email = data.Email;
                id = data.ID;
            } else {
                goto("/");
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

<div>
    <h1>User Profile</h1>
    <p><strong>Email:</strong> {email}</p>
    <p><strong>ID:</strong> {id}</p>
    <button class="buttonTest" type="button" on:click={logout}>Logout</button>
</div>
