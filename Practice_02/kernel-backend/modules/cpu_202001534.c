// This library provides functions and structures necessary to create and manage Linux kernel modules.
#include <linux/module.h>
/* This library provides functions and macros for printing messages to the kernel console, including 
macros like printk() and KERN_INFO.*/
#include <linux/kernel.h>
/* This library defines the module_init() and module_exit() macros to register the module's 
initialization and exit function, respectively.*/
#include <linux/init.h>
/* This library provides functions and structures for creating and managing files in the procfs file 
system, which provides a user interface for accessing kernel information and statistics.*/
#include <linux/proc_fs.h>
// This library provides functions to copy data between user space and kernel space.
#include <asm/uaccess.h>
/* This library provides functions and structures for writing data to stream files. Stream files are an 
efficient way to write large amounts of data to a file without having to store everything in memory.*/	
#include <linux/seq_file.h>
// This library provides structures and functions for working with kernel processes and tasks.
#include <linux/sched.h>
/* This library provides functions and structures for working with the kernel's virtual memory space, 
such as the mm_struct structure, which describes the current state of a process's virtual memory space.*/
#include <linux/mm.h>

// Module information
MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("CPU Information Module");
MODULE_AUTHOR("Erwin Fernando Vasquez Peñate");

/* "cpu" is a pointer to a task_struct that will be used to loop through all running processes on the 
system.*/
struct task_struct* cpu;
// "child" is a pointer to a task_struct that will be used to traverse the child processes of a process.
struct task_struct* child;
/* lstProcess is a pointer to a data structure that is used to traverse the list of child processes
of a process.*/
struct list_head* lstProcess;

//Function that will be executed every time the file is read with the CAT command
static int write_file(struct seq_file *the_file, void *v) {
    struct task_struct *task, *child_task;
    struct list_head *list;
    seq_printf(the_file, "[");
    for (task = &init_task; (task = next_task(task)) != &init_task;) {
        seq_printf(the_file, "{\"pid\":%d,\"name\":\"%s\",\"user\":%d,\"status\":%d",
            task_pid_nr(task), task->comm, task_uid(task), task->state);
        if (task->mm) {
            seq_printf(the_file, ",\"ram\":%ld", task_rss(task) << (PAGE_SHIFT - 20));
        }
        seq_printf(the_file, ",\"children\":[");
        list_for_each_entry(child_task, &task->children, sibling) {
            seq_printf(the_file, "%d,", task_pid_nr(child_task));
        }
        if (!list_empty(&task->children)) {
            the_file->buf[the_file->count-1] = ']';
        } else {
            seq_printf(the_file, "]");
        }
        seq_printf(the_file, "},");
    }
    the_file->buf[the_file->count-1] = ']';
    return 0;
}


//Function that will be executed every time the file is read with the CAT command
static int when_open(struct inode *inode, struct file *file){
    return single_open(file, write_file, NULL);
}
//If the kernel is 5.6 or higher, use the proc_ops structure
static struct proc_ops operations ={
    .proc_open = when_open,
    .proc_read = seq_read
};

//Function to execute when inserting the module in the kernel with insmod
static int _insert(void){
    proc_create("cpu_202001534", 0, NULL, &operations);
    printk(KERN_INFO "Erwin Fernando Vasquez Peñate\n");
    return 0;
}
//Function to execute when removing the kernel module with rmmod
static void _remove(void){
    remove_proc_entry("cpu_202001534", NULL);
    printk(KERN_INFO "Primer Semestre 2023\n");
}
module_init(_insert);
module_exit(_remove);