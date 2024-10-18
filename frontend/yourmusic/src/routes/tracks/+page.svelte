<script>
    let audio;
    const playAudio = () => {
       if(audio.paused){
        audio.play();
       } else{
        audio.pause();
       }
       
    }
    const songs = ["ena","duo","tria","tria","tria","tria","tria","tria","tria","tria"]
</script>
<main>
    <div class = "songs">
    {#each songs as song }
     <div class = "song">
        <img class = "profile-details" src="./image.png" alt="image of artist">
        <div class = "song-details" >
        <p>artist name</p>
        <p>{song}     00:00</p>
        
        </div>
        <div class = "song-details" >
        
            <button class="play-song" on:click={playAudio}>►</button>
            
            </div>

      </div>
    {/each}
    </div>
<div class="audio-player">
    <div class = "audio-player-details"> 
        <img class = "profile" src="./image.png" alt="image of artist">
        <div class  >
        <p>artist name</p>
        <p>song name</p>
        </div>
    </div>
    <div class = "audio-player-center">
    <audio bind:this={audio}  id="audio" src="./music.mp3"></audio>
    <div class="button-wrapper">
        <button class="audiobutton" on:click={playAudio}>►</button>
    </div>
    <div class="slider-wrapper">
        <div class="time-display">0:00</div>
        <input type="range" id="seek-bar" value="0" max="100" step="1">
        <div class="time-display">0:00</div>
    </div>
   
   
    </div>
</div>
</main>
<style>
    main{
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        width:  100%;
        height: 100%;
        background-color: rgb(0, 0, 0);
    }
    .profile-details{
        width: auto;
        height: 100%;
        margin-left: 25%;
        color: white;
    }

    .song-details{
        width: 20%;
        font-size: 3vh;
        color: white;
        letter-spacing: 0.2vh;
    }
    .play-song{
        font-size: 4vh;
    }
    .songs {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        flex-wrap: nowrap;   
        overflow-x: auto;
        align-items: center;
   
        width: 100%;
        height: 100%;
        background-color: rgb(12, 43, 71);
    }
    .song{
        width: 50%;
        height: 20%;
        display: flex;
        align-items: center;
        border-radius: 5vh;
        flex-direction: row;
        border: 1vh solid rgb(20, 12, 2);
        border-radius: 2vh;    
        margin: 1vh;
        padding: 1vh;
        background-color: rgb(20, 12, 2);
      
    }
    .profile{
        margin-top: 1vh;
        max-width: 5em;
        max-height: 5em;
    }
    .audio-player-details{
        display: flex;
        flex-direction: row;
        background-color: aqua;
        width: 25%;
        height: 100%;
    }
    .audio-player{
        display: flex;
        align-items: flex-end;
        justify-content: flex-start;
        width:  100%;
        height: 10%;
        background-color: rgb(36, 12, 247);
     
      
     
    }
    .button-wrapper{
        margin-top: 3vh;
    }
    .slider-wrapper {
        display: flex;
        align-items: center;
        gap: 10px; 
        width: 100%;
        justify-content: center;
        margin-top: 1vh;
    }

    .time-display {
        color: white;
        font-size: 14px;
    }
    .audio-player-center{
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        width: 55%;
        height: 100%;
        background-color: darkgreen;
     
    }

    input[type="range"]{
        width: 60%;
       
    

    }
    @media(max-width :1135px){
        .song-details{
        width: 10%;
        font-size: 2.5vh;
        color: white;
        letter-spacing: 0vh;
        }
        .profile-details{
      
        margin-left: 1vh;
    }
        .play-song{
        font-size: 3vh;
        margin-left: 3.5vh;
    }
    
    }
    @media(max-width :805px){
        .audio-player-details{
        width:  100%;
        height: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .audio-player{
  
       
        width:  100%;
        height: 50%;
        display: flex;
        align-items: flex-start;
        justify-content: center;
        flex-direction: column;
        background-color: rgb(25, 0, 250);
       
    }

    .slider-wrapper {
        width:  100%;
        height: 100%;
        display: flex;
        align-items: flex-end;
        justify-content: center;
    }
    .audio-player-center{
     
        width:  100%;
        height: 70%;
        display: flex;
        align-items:center;
        justify-content: center;
       
      
        background-color: rgb(7, 100, 7);
     
    }
    .time-display {
       margin-bottom: 5vh;
    }
    .button-wrapper{
        margin-top: 7vh;
    }
    input[type="range"]{
        width: 60%;
        margin-bottom: 5vh;

       
    

    }
    }
</style>