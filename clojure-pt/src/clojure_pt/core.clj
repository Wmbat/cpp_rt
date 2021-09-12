(ns clojure-pt.core
  (:gen-class))

(defn -main 
  [& args]
  (def image_width 256)
  (def image_height 256)
 
  (def rows (take image_width (range))) 
  (def cols (take image_height (range)))

  (for [row rows
        col cols]
    [row col]
      (print row) 
      (print ":")
      (print col)))
