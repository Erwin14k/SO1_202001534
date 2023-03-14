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
// Sysinfo alternative
#include <linux/hugetlb.h>


MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("RAM Information Module");
MODULE_AUTHOR("Erwin Fernando Vásquez Peñate");
//Function that is executed every time the file is read with the CAT command
static int write_file(struct seq_file *file, void *v){   
    struct sysinfo info;
    si_meminfo(&info);
    // Capture data in mb.
    // Total ram
    seq_printf(file, "{\"total_ram\":");
    seq_printf(file, "%ld", info.totalram* info.mem_unit / 1024/ 1024);
    // Free ram
    seq_printf(file, ",\"free_ram\":");
    seq_printf(file, "%ld", info.freeram* info.mem_unit / 1024 / 1024);
    // Occupied ram
    seq_printf(file, ",\"ram_occupied\":");
    seq_printf(file, "%ld", (info.totalram-info.freeram) * info.mem_unit / 1024 / 1024);
    seq_printf(file, "}\n");
    return 0;
}

//Function that is executed every time the file is read with the CAT command
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
    proc_create("ram_202001534", 0, NULL, &operations);
    printk(KERN_INFO "202001534\n");
    return 0;
}

//Function to execute when removing the kernel module with rmmod
static void _remove(void){
    remove_proc_entry("ram_202001534", NULL);
    printk(KERN_INFO "Sistemas Operativos 1 Seccion A\n");
}

module_init(_insert);
module_exit(_remove);