#include <pthread.h>
#include <stdio.h>

int i = 0 ; // The c language doesnt have a global keyword. However, variables declared outside a function have "file scope".

void* createThread_1()
{
    for ( int n = 0; n < 1000000 ; n++)
    {
        i++;
    }
    return NULL;
}

void* createThread_2()
{
    for ( int n = 0; n < 1000000 ; n++)
    {
        i--;
    }
    return NULL;
}

int main()
{
    
    pthread_t thread_1;
    pthread_create(&thread_1, NULL, createThread_1,NULL);
    pthread_t thread_2;
    pthread_create(&thread_2, NULL, createThread_2, NULL);

    pthread_join(thread_1, NULL);
    pthread_join(thread_2, NULL);

    printf("%d\n",i);

    return 0;
}

