<script lang="ts">
  import AddFriendButton from "./AddFriendButton.svelte";
</script>
<div class="container">
  <div><AddFriendButton /></div>
  <div><!-- empty --></div>
</div>
<style>
  .container {
    display: flex;
    justify-content: space-evenly;
    padding: 3px;
    box-sizing: border-box;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, .5);
  }
</style>
