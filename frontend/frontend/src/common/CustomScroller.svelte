<!-- CustomScroller.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  
  export let height;
  let wrapper: HTMLElement;
  let content: HTMLElement;
  let scrollThumb: HTMLElement;
  let isDragging = false;
  let startY = 0;
  let scrollTop = 0;

  function updateThumbPosition() {
    const contentHeight = content.scrollHeight;
    const viewportHeight = wrapper.clientHeight;
    const scrollPercent = scrollTop / (contentHeight - viewportHeight);
    const thumbHeight = Math.max(30, (viewportHeight / contentHeight) * viewportHeight);
    const maxThumbPos = viewportHeight - thumbHeight;
    
    scrollThumb.style.height = `${thumbHeight}px`;
    scrollThumb.style.transform = `translateY(${maxThumbPos * scrollPercent}px)`;
  }

  function onMouseDown(e: MouseEvent) {
    isDragging = true;
    startY = e.clientY - scrollThumb.getBoundingClientRect().top;
    document.addEventListener('mousemove', onMouseMove);
    document.addEventListener('mouseup', onMouseUp);
  }

  function onMouseMove(e: MouseEvent) {
    if (!isDragging) return;
    
    const contentHeight = content.scrollHeight;
    const viewportHeight = wrapper.clientHeight;
    const thumbHeight = Math.max(30, (viewportHeight / contentHeight) * viewportHeight);
    const maxThumbPos = viewportHeight - thumbHeight;
    
    const y = e.clientY - wrapper.getBoundingClientRect().top - startY;
    const scrollPercent = Math.max(0, Math.min(1, y / maxThumbPos));
    scrollTop = scrollPercent * (contentHeight - viewportHeight);
    content.style.transform = `translateY(-${scrollTop}px)`;
    updateThumbPosition();
  }

  function onMouseUp() {
    isDragging = false;
    document.removeEventListener('mousemove', onMouseMove);
    document.removeEventListener('mouseup', onMouseUp);
  }

  function onWheel(e: WheelEvent) {
    e.preventDefault();
    const contentHeight = content.scrollHeight;
    const viewportHeight = wrapper.clientHeight;
    
    scrollTop = Math.max(0, Math.min(contentHeight - viewportHeight, scrollTop + e.deltaY));
    content.style.transform = `translateY(-${scrollTop}px)`;
    updateThumbPosition();
  }

  onMount(() => {
    updateThumbPosition();
    wrapper.addEventListener('wheel', onWheel, { passive: false });
    
    return () => {
      wrapper.removeEventListener('wheel', onWheel);
    };
  });
</script>

<div class="custom-scroll" bind:this={wrapper} style="height: {height}">
  <div class="content" bind:this={content}>
    <slot />
  </div>
  <div 
    class="scrollbar"
    on:mousedown|stopPropagation
  >
    <div
      class="thumb"
      bind:this={scrollThumb}
      on:mousedown={onMouseDown}
    />
  </div>
</div>

<style>
  .custom-scroll {
    position: relative;
    width: 100%;
  }

  .content {
    position: relative;
    width: calc(100% - 12px);
  }

  .scrollbar {
    position: absolute;
    right: 0;
    top: 0;
    width: 8px;
    height: 100%;
    background: rgba(0, 0, 0, 0.1);
    border-radius: 4px;
  }

  .thumb {
    width: 100%;
    background: rgba(0, 0, 0, 0.4);
    border-radius: 4px;
    cursor: pointer;
    transition: background 0.2s;
  }

  .thumb:hover {
    background: rgba(0, 0, 0, 0.6);
  }
</style>
