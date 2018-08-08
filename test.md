## Asset-Manager high level design

Asset manager is responsible for the managing the DCP/KDM/SPL on the SMS. It is divided in to different components as mentioned below.

<iframe frameborder="0" style="width:100%;height:1058px;" src="https://www.draw.io/?lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1&title=asset_manger_block_diagram.xml#R7Vxbb6M4GP01kWYfinw3fpx2LrvSdFWpK%2B3MI01ogoaELJC22V%2B%2FJmDAl0wIIQkz2vJQMNjg892OP9uZ4Lvl2%2Bc0WC%2Fuk1kYTxCYvU3whwlCkCAm%2FxUl27LER35ZME%2BjWfVQU%2FAY%2FRtWhaAq3USzMNMezJMkzqO1XjhNVqtwmmtlQZomr%2Fpjz0msv3UdzEOr4HEaxHbp39EsX1S9oKAp%2Fz2M5gv1ZgiqO0%2FB9Ps8TTar6n0ThJ93f%2BXtZaDaqp7PFsEseW0V4Y8TfJcmSV6eLd%2FuwrjAVsFW1vu052793Wm4yrtUIKSs8RLEm1B98u7D8q0CY9edsKgAJvj2dRHl4eM6mBZ3X6X4ZdkiX8byCsrTWZRKgUTJSl6vkrTo7q39UdV3voRpHr61iqqP%2FBwmyzBPt%2FIRdVd4tKxTaRRWiL828uG8Klu0ZMNpVTGodGJeN97gIk8qaNwwIXEYpkLs6%2B5drZU3eFItgB9CoJSvAoAzGwDGHABIqE4HAGILgPdZFuay6DFP0sKcEAuWhTKsnrL1Tujg3adIdk0%2Bsc3ycPmb%2FYR1Hctvvn1K5dm8OPt6%2F0VWL1rJ7JuWAKQtrYvTWZAHmfyo8LC22uKqJX1YNbe18VuCQL5DEBANIAjC9wriPlhJOaQF4GH6EskumwjJHuU6AFmeJt%2FDuyRO0p3BrgrQnqM4NoqCOJoXJj2VQMlX4NsCn0g6zPfVjWU0m8X7ENc9yADuAOu2gB0igC5bIANIANkSsIAOV7P3RRgqEIuDLIumO6yDNLeLW9II36L8a4GSdHbl1bdDmGXJJp2qqFUZqXzPPFT6XD0XzrSIZyPbgo46kFNlaRgHefSix0kXnNUbHpJolTeC475HGRaYSIogbY1CTY4QUOIRvc2yg1Uz7UhmtAwB8SCWIYAI6AMs32Q0DXyz6RInq%2BmdAtTIdNMJpcG2Vd4F00XhBE8LqwPYDDfih3DED%2BKS%2FADxk%2FkWPn8mefQsPciOK5R6tbUjwl7M4EUwI9g3eAd1uRpCzwMbRz1djdOnCKp5FQ8gUl0%2FhGkkP65w7DtllICl27ISBlgVFLWgB6CvCsxqnZwUE7aTUuoxEicFdUuB0GRQXX0SMxoCvtHQcB7ItwnaYFGp0QYiNVLXhlKHXNrQKB4GfkvxboDSIYfmdVKhekjYViHxvwqdqkLssAqde5BjdJc5RnmEQFtsEOEBopRrlGcPO85V8leQfW%2FT%2BGH5Qi3dw8HvImBz22EVAMAr9ru6KzwBGEWQYyHNlmCdADjiP2ceE7Sp4mADbICBB9%2Bvntk6WLXVCKsbss32veuDCwXxJP3nvoTEL1BjI0FX%2BdUO6KLxoguox3mFbIEYGQu6tu5%2BfJEdzWTFL8lc3nkXPe%2FC7j%2BbKA1nv%2F0Ay8tQf2aEcCfxRw4GMQRcwqWMPYm%2FOi9JP6cao2uzuTLrsBc6jcRXhL3NwPxxMTCK9NwCFT0ZGOVGQ2YSaTgGJjqM986eZjYz7ZSpgnaCEzBbctCEuJfq26mUrhjU8PUH4SJdpB3GaieGDtHd312myx3mmfZJdTjVRrpid8zk1An%2BkwCgDgAuN7a4D7MsmEereYd5lBMjbS3q7qzFKyIpYIALIMcFlBqM23Y%2FUk4eIYAjQXyMJJfk5xKbPSY2EpePm6dsmkZPPxyxXQhILrk1IgBKyHxAiRG4kI0jw3KsQzmngFPhc5%2FYMA5BZ6hj0urhj0mRHI%2Fj7Pq4EU%2FTJTXHdn3cbBpYTStYc63N5N80jiQUN%2FZAeqyzf0fEquruDTU4CoQOjlIHLn02%2FHSxQNCfowwWzDAHVyZqHVjMPhDA6TBcpo8n0Bb8k3TRRUwGJqPdHfJlutw%2F0b23Kz8VGRWuJQy%2FJBkVx6fQxktGxcFZ9POR0eOBHCsZFXYq8oxk9HjcxkpG6zWufdio7WBHy0aPHj1cnY2eQMSGimYMeAKy1p8%2B8YAJ8AhtHcTCh0MHPpAOoLjYReJ%2ByWjXqMJRTOuM0Nu%2B1ohYD5unOMoWwwesY7BQQyoqPJ8AoQ7dqLHHcOtQWxxaOGJp%2BIK1Docn9qHnw%2FZxOsjEZvBnDGg9cEXMw5SB5kDHIot85kHaOi4CLOswTui%2B%2FA71WtOrLKg906aSkiOZaVOToM16%2BJ4zbcYKVQhMFzTcTBuzuUw30R61XK6ea%2F2235xskSt71kTOxyVybO3HQcZ6165CR77ZFOJGU8OJHYIOCZ2DJt0ImGFdxB5AdUGvpY0ucxfjXh2LzKFt36WN0BylDCn3DlmuM9h74%2F05RNo6DMDVdS89cfkINbAdiZ7UeX3lIUxv3lVPsLnbx3Q1Ay6B7bu3Z7hNPK6lNXhckiXGVkRqbmzrKlniE8P3E1HPapxBuh2mbU7euSWZ7sRY7M5OsXTXfgll%2FSPRh59xsTuEg5o6N229FRX4HhpoSK%2Fftgf%2FmoKv8zhqm39fKkCF3hA%2FIxWAdhbdEnwrXdY5N%2BaQ3%2F5Mg9A13VeKfnAP8RDzCBA6UuCbtyiOgnR7U2yfPn6r9ofb%2B8dz7tF2iEGDGx42BDegwJdcDEBOCfapYMQfAF%2BlEFfYInysU6l1wd40PJp8Qt9wYuUTBgsn8rL5jY7y8eaHUPDH%2FwA%3D"></iframe>

### Storage
It store/retrieve the asset manger data on the disk. At high level data is devided in two types.

1. **Assets**  
KDM xml files, SPL xml files and DCPs are the assets. These xml files are maintained in a directory structure as defined below.
   * **kdm** :  Directory to keep kdm xml files on the SMS
   * **spl** : Directory to keep spl xml files on the SMS
   * **dcp** : Directory to keep DCPs on the SMS. Under this directory,  several directories will be created for each DCP, one directory will map to one DCP.   These directories keeps the PKL, CPL, asset-map and mxf files.	            
	      
            

2. **Auxiliary-data**  
      Data which need to  persist across several life cycles of the service. This information would be generated during processing different request. Any DBMS(like sqlite) can be used to maintain this information.
	 

### Asset cache
1. It is introduced in the system to enable faster access to the Assets stored on the disk. This will provide the APIs to store/retrieve the assets and auxiliary-data to/from system.
2. To avoid rebuilding the new cache at every start of service, This cache is persisted on the disk using DBMS(like sqlite).
3. Cache can be rebuilt at any point from the available data in the storage. Cache rebuild might be required in scenarios like software upgrade where older cache format is changed. 
4. If auxiliary-data is corrupted or becomes invalid, then it is ignored during cache rebuild (loss of auxiliary data).
		
### Task manager
1. Task manager gets/stores the assets and auxiliary-data using asset cache.
2. Task manager provides
   1. Interface(APIs) for asset management tasks/operations.
       * Ingest of CPL/SPL/DCP. 
       * Deletion of CPL/SPL/DCP.
       * Get the status of the ongoing asset management tasks/operations.
       * List the Assets and details of particular asset.	
   2. Interface(APIs) for controlling asset management operations (stop the allowing new ingest , etc)   	
   3. Auto delete functionality for CPL/SPL/DCP according to the business rules.
  

### Notification relay
Different components of the asset manager (like task manager, asset cache) will call notification relay API to generate the notifications. These notification will be forwarded the the registered handler. Notifications handler would be

1. Asset Manager Messaging.
2. Logs/Events writer module.

### Messaging:

Asset manager communicates with the external world with messaging layer. It  have follwing components
     
1. **Message/API call/Request-Response** : Handle the messages from external world. When ever it receives any message, it will call  appropriate APIs of the task manager. Depending on the API data, it will construct the response message and send it back to the peer. 
2. **Notification** :
   Mechanism to publish the notifications to external world. Who ever want to get the notification need the subscribe it. It will receive notification from notification relay and forward it to subscriber.
     
