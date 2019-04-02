 /*************************************************************************
    > File Name: 4_2.thread.c
    > Author: Shen Bo
    > mail: nichol_shen@yahoo.com
    > Created Time: Fri Mar  8 17:33:26 2019
 ************************************************************************/

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

void *count();
pthread_mutex_t mutex1 = PTHREAD_MUTEX_INITIALIZER;
int counter = 0;

void add()
{
    printf("add function called\n");
}

void main()
{
    int rc1, rc2;
    pthread_t thread1, thread2;

    // Create thread, function would be executed indevidully in every thread

    if (rc1 = pthread_create(&thread1, NULL, &add, NULL))
    {
        printf("Thread1 creation failed: %d\n", rc1);
    }

    if (rc1 = pthread_create(&thread2, NULL, &add, NULL))
    {
        printf("Thread2 creation failed: %d\n", rc2);
    }

    // Wait all threads executed
    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    exit(0);
}

void *count()
{
    pthread_mutex_lock(&mutex1);
    counter ++;
    printf("Counter value : %d\n", counter);
    pthread_mutex_unlock(&mutex1);
}
