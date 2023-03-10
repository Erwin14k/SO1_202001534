#include <linux/module.h>
// Header to use KERN_INFO
#include <linux/kernel.h>
//Header to use module_init
#include <linux/init.h>
//Header to use proc_fs
#include <linux/proc_fs.h>
#include <asm/uaccess.h>	
#include <linux/seq_file.h>
#include <linux/sched.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("CPU Information Module");
MODULE_AUTHOR("Erwin Fernando Vasquez Peñate");

struct task_struct* cpu;
struct task_struct* child;
struct list_head* lstProcess;

//Function that will be executed every time the file is read with the CAT command
static int write_file(struct seq_file *the_file, void *v){   
    int ram, split, child_split;
    split = 0;
    child_split = 0;
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
        seq_printf(the_file, "]}");
        split = 1;
    }
    seq_printf(the_file, "]");
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