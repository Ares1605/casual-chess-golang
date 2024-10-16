<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { SignIn } from '../wailsjs/go/main/App.js';
  import { googleuser } from 'wailsjs/go/models';

  const dispatch = createEventDispatcher<{
    signIn: googleuser.GoogleUser;
  }>();

  function signIn() {
    SignIn().then(user => {
      dispatch('signIn', user)
    })
  }
</script>
<img class="title" src="/logos/casual-chess.png" alt="casual chess">
<div>
  <img class="board" src="/logos/chess-board.png" alt="board">
  <button on:click={signIn}>Sign In</button>
</div>

<style>
  .title {
    display: block;
    margin: auto;
    height: 100px;
  }
  div {
    position: relative;
    display: block;
    margin: auto;
    width: 600px;
    height: 600px;
  }
  .board {
    width: 100%;
    height: 100%;
  }
  button {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-size: 25px;
    border: 4px solid #42210d;
    width: 175px;
    height: 50px;
    border-radius: 25px;
    background-color: #d4bea1;
    color: #42210d;
    cursor: pointer;

    &:active {
      background-color: #42210d;
      color: #d4bea1;
    }
  }
</style>
