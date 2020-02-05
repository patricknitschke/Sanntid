#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mutex_for_i;

// Note the return type: void*
void* incrementingThreadFunction(){
    pthread_mutex_lock(&mutex_for_i);
    for (int j = 0; j < 1000001; j++) {
	// TODO: sync access to i
	i++;
    }
    pthread_mutex_unlock(&mutex_for_i);
    return NULL;
}

void* decrementingThreadFunction(){
    pthread_mutex_lock(&mutex_for_i);
    for (int j = 0; j < 1000003; j++) {
	// TODO: sync access to i
	i--;
    }
    pthread_mutex_unlock(&mutex_for_i);
    return NULL;
}


int main(){
    pthread_t incrementingThread, decrementingThread;
    pthread_mutex_init(&mutex_for_i, NULL);

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    pthread_mutex_destroy(&mutex_for_i);

    
    printf("The magic number is: %d\n", i);
    return 0;
}
