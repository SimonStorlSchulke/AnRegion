<script lang="ts">
  import {CompIt, GetProgress} from '../wailsjs/go/main/App.js'

  let progress: number = 0
  let folderRegion: string = ""
  let folderOriginal: string = ""
  let ignorePrefixes: string = ".bumpNormals;.VRayDenoiser"
  let running = false
  let overwrite: boolean = false
  const imageMagickPath = "composite.exe"

  function comp(): void {
    progress = 0
    CompIt(imageMagickPath, folderRegion, folderOriginal, ignorePrefixes, overwrite).then()
  }

  window.setInterval(() => {
    GetProgress().then(val => {
      progress = val
      console.log(progress)
    })
  }, 300)

</script>

<main>
  <div>
    <h2>AnRegion</h2>
  </div>
  <div class="input-box" id="input">
    <span style="display: inline-block; min-width: 200px;" class="result" id="result">Original Renderings Folder</span>
    <input style="width: 500px;" autocomplete="off" bind:value={folderOriginal} class="input" id="name" type="text"/><br>
    <span style="display: inline-block; min-width: 200px;" class="result" id="result">Region Renderings Folder</span>
    <input style="width: 500px;" autocomplete="off" bind:value={folderRegion} class="input" id="name" type="text"/><br><br>

    <span style="display: inline-block; min-width: 200px;" class="result" id="result">Ignore filenames including:</span>
    <input style="width: 500px;" autocomplete="off" bind:value={ignorePrefixes} class="input" id="name" type="text"/><br><br>

    <label style="display: inline-block; min-width: 200px;" for="horns">Overwrite original renders*</label>
    <div style="width: 500px; display: inline-block; text-align: left;"><input bind:checked={overwrite} type="checkbox" id="horns" name="horns"> </div><br><br>

    {#if progress != 0 && progress != 100}
    <div id="progress" style="background-image: linear-gradient(90deg, #4288e6 {progress}%, black {progress+0.01}%);">Progress</div>
    {:else}
    <button style="width: 300px" class="btn" on:click={comp}>Comp it</button>
    {/if}
    <br><br><span style="color: #fff7;">*if this is not checked, the composited files will instead be saved in a subfolder of the Original Renderings Folder /anrender</span>
  </div>

</main>

<style>


  .result {
    height: 20px;
    line-height: 40px;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  #progress {
    background: #333;
    border-radius: 13px;
    height: 20px;
    width: 300px;
    padding: 3px;
    display: block;
    margin: auto;
    margin-top: 0px;
    color: rgba(255, 255, 255, 0.735);
}


</style>
