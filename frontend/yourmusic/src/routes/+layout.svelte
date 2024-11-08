<script>
import '../styles/global.css'
let { children } = $props();
</script>

<header>

</header>
<main>
{@render children()}
</main>
<footer>

</footer>
<style>

    main{
        width: 100%;
        height: 100%;
        background-color: black;
    }
</style>
