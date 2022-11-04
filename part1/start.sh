fastqc e_coli_fastq.gz
mv e_coli_fastqc.html e_coli.html
rm e_coli_fastqc.zip
minimap2 -d ref.mmi ref_e_coli.fna.gz
minimap2 -a ref.mmi e_coli_fastq.gz > alignment.sam
samtools flagstat alignment.sam
