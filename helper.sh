Help()
{
   # Display Help
   echo "helper scripts"
   echo
   echo "syntax: helper -[e|u|h]"
   echo "options:"
   echo "e     export .env"
   echo "u    unset .env"
   echo "h     help"
   echo
}

while getopts ":h:e:u" option; do
   case $option in
      h) # display Help
         Help
         exit;;
      e) # display Help
         export $(grep -v '^\#' .env | xargs -d '\n')
         echo ".env exported"
         exit;;
      u) # display Help
         unset $(grep -v '^\#' .env | sed -E '/AF|PO/p' | xargs)
         echo ".env unsetted"
         exit;;
     \?) # Invalid option
         echo "Error: Invalid option"
         exit;;
   esac
done
