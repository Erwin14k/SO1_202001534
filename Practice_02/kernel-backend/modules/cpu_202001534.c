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
static int cpu_percentage(void);

//Function that will be executed every time the file is read with the CAT command
static int write_file(struct seq_file *the_file, void *v){   
    int ram, split, child_split;
    split = 0;
    child_split = 0;
    int percentage;
    percentage=cpu_percentage();
    seq_printf(the_file, "[");
    seq_printf(the_file, "%ld",percentage);
    seq_printf(the_file, "],\n");
    seq_printf(the_file, "[");
    for_each_process(cpu){
        if(split){
            seq_printf(the_file, ",");
        }
        seq_printf(the_file, "{\"pid\":");
        seq_printf(the_file, "%d", cpu->pid);
        seq_printf(the_file, ",\"name\":");
        seq_printf(the_file, "\"%s\"", cpu->comm);
        seq_printf(the_file, ",\"user\":");
        seq_printf(the_file, "%d", cpu->real_cred->uid);
        seq_printf(the_file, ",\"status\":");
        seq_printf(the_file, "%d", cpu->__state);
        if (cpu->mm) {
            ram = (get_mm_rss(cpu->mm)<<PAGE_SHIFT)/(1024*1024);
            seq_printf(the_file, ",\"ram\":");
            seq_printf(the_file, "%d", ram);
        }
        seq_printf(the_file, ",\"children\":[");
        child_split = 0;
        list_for_each(lstProcess, &(cpu->children)){
            child = list_entry(lstProcess, struct task_struct, sibling);
            if(child_split){
                seq_printf(the_file, ",");
            }
            seq_printf(the_file, "%d", child->pid);
            child_split = 1;
        }
        seq_printf(the_file, "]}\n");
        split = 1;
    }
    seq_printf(the_file, "]\n");
    return 0;
}

// Function to calculate the cpu percentage
static int cpu_percentage(void){
    struct file *the_file;
    char lecture[256];
    int user,niced,system,idle,iowait,irq,suaveirq,steal,guest,guest_nice;
    int total;
    int percentage;
    the_file=filp_open("/proc/stat",O_RDONLY,0);
    memset(lecture,0,sizeof(lecture));
    kernel_read(the_file,lecture,sizeof(lecture)-1,&the_file->f_pos);
    sscanf(lecture,"cpu %d %d %d %d %d %d %d %d %d %d",
        &user,&niced,&system,&idle,&iowait,&irq,&suaveirq,&steal,&guest,&guest_nice);
    total=user+niced+system+idle+iowait+irq+suaveirq+steal+guest+guest_nice;
    percentage=100-(idle*100/total);
    filp_close(the_file,NULL);
    return percentage;    
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