// This a comment
/* Another comment */
/** Yet another comment */

// Create dest - leadingComment
try {
    fs.mkdirSync(dest, parseInt('0777', 8));
} catch (e) {
    // like Unix's cp, keep going even if we can't create dest dir - innerComment
}
