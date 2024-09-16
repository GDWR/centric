import { ref } from "vue";
import { defineStore } from "pinia";
import { useEnvironmentStore } from "./environment";

export const useImagesStore = defineStore("images", {
  state: () => {
    let images = ref([]);
    
    const { selectedEnvironmentId } = useEnvironmentStore()
    fetch(`/api/v1/environments/${selectedEnvironmentId}/images`)
      .then(async response => {
        const data = await response.json();
        images.value = data;
      });

    return { images }
  },

  getters: {
    count: (state) => state.images.length,
  },
});
