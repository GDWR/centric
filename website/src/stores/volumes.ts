import { ref } from "vue";
import { defineStore } from "pinia";
import { useEnvironmentStore } from "./environment";

export const useVolumesStore = defineStore("volumes", {
  state: () => {
    let volumes = ref([]);
    
    const { selectedEnvironmentId } = useEnvironmentStore()
    fetch(`/api/v1/environments/${selectedEnvironmentId}/volumes`)
      .then(async response => {
        const data = await response.json();
        volumes.value = data;
      });

    return { volumes}
  },

  getters: {
    count: (state) => state.volumes.length,
  },
});
