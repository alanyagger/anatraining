#include <stddef.h>
#include "ex10_ll_cycle.h"

int ll_has_cycle(node *head) {
    /* TODO: Implement ll_has_cycle */
    node* fast_ptr=head;
    node* slow_ptr=head;
    node* temp_ptr=head;
    if (head==NULL)
    {
    	return 0;
	}
    do
    {
    	temp_ptr=(*fast_ptr).next;
    	if (temp_ptr==NULL)
    	{
    		return 0;
		}
		fast_ptr=(*temp_ptr).next;
		if (fast_ptr==NULL)
    	{
    		return 0;
		}
		slow_ptr=(*slow_ptr).next;
	}while(fast_ptr!=slow_ptr);
	return 1;
} 
