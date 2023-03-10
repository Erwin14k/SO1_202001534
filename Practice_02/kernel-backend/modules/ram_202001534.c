#include <linux/module.h>
// Header KERN_INFO
#include <linux/kernel.h>

//Header module_init & module_exit
#include <linux/init.h>
//Header proc_fs
#include <linux/proc_fs.h>
#include <asm/uaccess.h>	
// seq_file library
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
    seq_printf(file, "{\"totalram\":");
    seq_printf(file, "%ld", info.totalram);
    seq_printf(file, ",\"freeram\":");
    seq_printf(file, "%ld", info.freeram);
    seq_printf(file, "}");
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