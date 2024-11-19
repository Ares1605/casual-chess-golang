<script lang="ts">
  export let user: any;

  let opacity = 0;

  setTimeout(() => opacity = 1);
</script>
<div style:opacity={opacity} class="container">
  <img src={user.profile_url} alt="profile">
  {user.username}
</div>
<style>
  img {
    height: 50px;
  }
  .container {
    z-index: 2;
    transition: .2s opacity;
    position: absolute;
    width: 150px;
    background-color: var(--color-primary);
    border: 3px solid var(--color-primary-darker);
    color: var(--color-neutral);
    border-radius: 3px;
    padding: 10px;
    transform: translateX(-50%);
  }
</style>
